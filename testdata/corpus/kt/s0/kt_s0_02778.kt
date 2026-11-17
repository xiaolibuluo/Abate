package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22920<A, B>(val first: A, val second: B)

fun combine22920(left: Int, right: Int): Pair22920<Int, Int> {
    return Pair22920(left + 22, right + 10)
}
data class Pair22921<A, B>(val first: A, val second: B)

fun combine22921(left: Int, right: Int): Pair22921<Int, Int> {
    return Pair22921(left + 29, right + 21)
}
