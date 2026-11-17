package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22880<A, B>(val first: A, val second: B)

fun combine22880(left: Int, right: Int): Pair22880<Int, Int> {
    return Pair22880(left + 67, right + 43)
}
