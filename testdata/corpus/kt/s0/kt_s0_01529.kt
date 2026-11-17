package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair12670<A, B>(val first: A, val second: B)

fun combine12670(left: Int, right: Int): Pair12670<Int, Int> {
    return Pair12670(left + 83, right + 33)
}
object Registry12671 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12671 = "registry-90"
    }
}
