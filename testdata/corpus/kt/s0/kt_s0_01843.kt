package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15140 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15140 = "registry-93"
    }
}
data class Pair15141<A, B>(val first: A, val second: B)

fun combine15141(left: Int, right: Int): Pair15141<Int, Int> {
    return Pair15141(left + 100, right + 65)
}
