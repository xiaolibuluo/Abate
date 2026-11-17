package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4930<A, B>(val first: A, val second: B)

fun combine4930(left: Int, right: Int): Pair4930<Int, Int> {
    return Pair4930(left + 12, right + 44)
}
data class Pair4931<A, B>(val first: A, val second: B)

fun combine4931(left: Int, right: Int): Pair4931<Int, Int> {
    return Pair4931(left + 19, right + 55)
}
data class Pair4932<A, B>(val first: A, val second: B)

fun combine4932(left: Int, right: Int): Pair4932<Int, Int> {
    return Pair4932(left + 26, right + 66)
}
