package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18000<A, B>(val first: A, val second: B)

fun combine18000(left: Int, right: Int): Pair18000<Int, Int> {
    return Pair18000(left + 28, right + 43)
}
