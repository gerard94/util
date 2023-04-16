/*
 float: Mathematical coprocessor analyzer.

Copyright (C) 2004...2020 Gérard Meunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
*/

package float
	
	/*
		Ref: W. H. Press, S. A. Teukolsky, W. T. Vetterling, B. P. Flannery, Numerical Recipes in Fortran 77, second edition, ch. 20, pp. 881-886 (Diagnosing Machine Parameters); http://www.nr.com/.

		-	Ibeta is the radix in which numbers are represented, almost always 2, but occasionally 16, or even 10.
		-	It is the number of base-Ibeta digits in the floating-point mantissa M.
		-	Machep is the exponent of the smallest (most negative) power of Ibeta that, added to 1.0, gives something different from 1.0.
		-	Eps is the floating-point number ibetamachep, loosely referred to as the “floating-point precision.”
		-	Negep is the exponent of the smallest power of Ibeta that, subtracted
		from 1.0, gives something different from 1.0.
		-	Epsneg is ibetanegep, another way of defining floating-point precision. Not infrequently Epsneg is 0.5 times Eps; occasionally Eps and Epsneg are equal.
		-	Iexp is the number of bits in the exponent (including its sign or bias).
		-	Minexp is the smallest (most negative) power of Ibeta consistent with there being no leading zeros in the mantissa.
		-	Xmin is the floating-point number ibetaminexp, generally the smallest (in magnitude) useable floating value.
		-	Maxexp is the smallest (positive) power of Ibeta that causes overflow.
		-	Xmax is (1- Epsneg) × ibetamaxexp, generally the largest (in magnitude) useable floating value.
		-	Irnd returns a code in the range 0...5, giving information on what kind of rounding is done in addition, and on how underflow is handled. See below.
		-	Ngrd is the number of “guard digits” used when truncating the product of two mantissas to fit the representation.

		The parameter Irnd needs some additional explanation. In the IEEE standard, bit patterns correspond to exact, “representable” numbers. The specified method for rounding an addition is to add two representable numbers “exactly,” and then round the sum to the closest representable number. If the sum is precisely halfway between two representable numbers, it should be rounded to the even one (low-order bit zero). The same behavior should hold for all the other arithmetic operations, that is, they should be done in a manner equivalent to infinite precision, and then rounded to the closest representable number.

		If Irnd returns 2 or 5, then your computer is compliant with this standard. If it returns 1 or 4, then it is doing some kind of rounding, but not the IEEE standard. If Irnd returns 0 or 3, then it is truncating the result, not rounding it — not desirable.

		The other issue addressed by Irnd concerns underflow. If a floating value is less than Xmin, many computers underflow its value to zero. Values Irnd = 0, 1, or 2 indicate this behavior. The IEEE standard specifies a more graceful kind of underflow: As a value becomes smaller than Xmin, its exponent is frozen at the smallest allowed value, while its mantissa is decreased, acquiring leading zeros and “gracefully” losing precision. This is indicated by Irnd = 3, 4 or 5.
	*/
	
	
	
	import (
		
		M	"git.duniter.org/gerard94/util/misc"
		
			"fmt"
	
	)
	
	var (
		
		Ibeta, It, Irnd, Ngrd, Machep, Negep, Iexp, Minexp, Maxexp int64
		Eps, Epsneg, Xmin, Xmax float64
	)
	
	func machar () {
		
		var (
			
			i, itemp, iz, j, k, mx, nxres int64
			a, b, beta, betah, betain, one, t, temp, temp1, tempa, two, y, z, zero float64
		
		)
		
		one = 1
		two = one + one
		zero = one - one
		
		// Determine Ibeta and beta by the method of M. Malcolm.
		a = one
		for {
			a = a + a
			temp = a + one
			temp1 = temp - a
			if temp1 - one != zero {break}
		}
		b = one
		for {
			b = b + b
			temp = a + b
			itemp = int64(temp - a)
			if itemp != 0 {break}
		}
		Ibeta = itemp
		beta = float64(Ibeta)
		
		// Determine It and Irnd.
		It = 0
		b = one
		for {
			It++
			b = b * beta
			temp = b + one
			temp1 = temp - b
			if temp1 - one != zero {break}
		}
		Irnd = 0
		betah = beta / two
		temp = a + betah
		if temp - a != zero {
			Irnd = 1
		}
		tempa = a + beta
		temp = tempa + betah
		if (Irnd == 0) && (temp - tempa != zero) {
			Irnd = 2
		}
		
		// Determine Negep and Epsneg.
		Negep = It + 3
		betain = one / beta
		a = one
		for i = 1; i <= Negep; i++ {
			a = a * betain
		}
		b = a
		for {
			temp = one - a
			if temp - one != zero {break}
			a = a * beta
			Negep = Negep - 1
		}
		Negep = - Negep
		Epsneg = a
		
		// Determine Machep and Eps.
		Machep = - It - 3
		a = b
		for {
			temp = one + a
			if temp - one != zero {break}
			a = a * beta
			Machep++
		}
		Eps = a
		
		// Determine Ngrd.
		Ngrd = 0
		temp = one + Eps
		if (Irnd == 0) && (temp * one - one != zero) {
			Ngrd = 1
		}
		
		// Determine Iexp.
		i = 0
		k = 1
		z = betain
		t = one + Eps
		nxres = 0
		for { // Loop until an underflow occurs, then exit.
			y = z
			z = y * y
			a = z * one // Check here for the underflow.
			temp = z * t
			if a + a == zero || M.AbsF64(z) >= y {break}
			temp1 = temp * betain
			if temp1 * beta == z {break}
			i++
			k = k + k
		}
		if Ibeta != 10 {
			Iexp = i + 1
			mx = k + k
		} else {	// For decimal machines only.
			Iexp = 2
			iz = Ibeta
			for k >= iz {
				iz = iz * Ibeta
				Iexp++
			}
			mx = iz + iz - 1
		}
		
		// To determine Minexp and Xmin, loop until an underflow occurs, then exit.
		for {
			Xmin = y
			y = y * betain
			a = y * one // Check here for the underflow.
			temp = y * t
			if a + a == zero || M.AbsF64(y) >= Xmin {break}
			k++
			temp1 = temp * betain
			if temp1 * beta == y && temp != y {
				nxres = 3
				Xmin = y
				break
			}
		}
		
		// Determine Maxexp, Xmax.
		Minexp = - k
		if mx <= k + k - 3 && Ibeta != 10 {
			mx = mx + mx
			Iexp++
		}
		Maxexp = mx + Minexp
		Irnd = Irnd + nxres // Adjust Irnd to reflect partial underflow.
		if Irnd >= 2 { // Adjust for IEEE-style machines.
			Maxexp = Maxexp - 2
		}
		i = Maxexp + Minexp
		// Adjust for machines with implicit leading bit in binary mantissa, and machines with radix point at extreme right of mantissa.
		if Ibeta == 2 && i == 0 {
			Maxexp = Maxexp - 1
		}
		if i > 20 {
			Maxexp = Maxexp - 1
		}
		if a != y {
			Maxexp = Maxexp - 2
		}
		Xmax = one - Epsneg
		if Xmax * one != Xmax {
			Xmax = one - beta * Epsneg
		}
		Xmax = Xmax / (beta * beta * beta * Xmin)
		i = Maxexp + Minexp + 3
		for j = 1; j <= i; j++ {
			if Ibeta == 2 {
				Xmax = Xmax + Xmax
			} else {
				Xmax = Xmax * beta
			}
		}
	} //machar
	
	func ShowPrecision () {
		fmt.Println()
		fmt.Println("Ibeta =", Ibeta)
		fmt.Println("It =", It)
		fmt.Println("Machep =", Machep)
		fmt.Println("Eps =", Eps)
		fmt.Println("Negep =", Negep)
		fmt.Println("Epsneg =", Epsneg)
		fmt.Println("Iexp =", Iexp)
		fmt.Println("Minexp =", Minexp)
		fmt.Println("Xmin =", Xmin)
		fmt.Println("Maxexp =", Maxexp)
		fmt.Println("Xmax =", Xmax)
		fmt.Println("Irnd =", Irnd)
		fmt.Println("Ngrd =", Ngrd)
		fmt.Println()
	} //ShowPrecision
	
	func init () {
		machar()
	} //init
