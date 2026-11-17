package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18980<A, B>(val first: A, val second: B)

fun combine18980(left: Int, right: Int): Pair18980<Int, Int> {
    return Pair18980(left + 41, right + 58)
}
