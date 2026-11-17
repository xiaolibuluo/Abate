package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3910<A, B>(val first: A, val second: B)

fun combine3910(left: Int, right: Int): Pair3910<Int, Int> {
    return Pair3910(left + 44, right + 62)
}
data class Pair3911<A, B>(val first: A, val second: B)

fun combine3911(left: Int, right: Int): Pair3911<Int, Int> {
    return Pair3911(left + 51, right + 12)
}
