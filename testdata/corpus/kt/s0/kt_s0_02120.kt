package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair17520<A, B>(val first: A, val second: B)

fun combine17520(left: Int, right: Int): Pair17520<Int, Int> {
    return Pair17520(left + 83, right + 12)
}
