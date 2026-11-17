package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry6350 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6350 = "registry-15"
    }
}
data class Pair6351<A, B>(val first: A, val second: B)

fun combine6351(left: Int, right: Int): Pair6351<Int, Int> {
    return Pair6351(left + 22, right + 12)
}
