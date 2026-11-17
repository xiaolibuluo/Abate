package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair12700<A, B>(val first: A, val second: B)

fun combine12700(left: Int, right: Int): Pair12700<Int, Int> {
    return Pair12700(left + 25, right + 54)
}
data class Pair12701<A, B>(val first: A, val second: B)

fun combine12701(left: Int, right: Int): Pair12701<Int, Int> {
    return Pair12701(left + 32, right + 65)
}
