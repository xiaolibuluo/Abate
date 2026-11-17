package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10740 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10740 = "registry-96"
    }
}
data class Pair10741<A, B>(val first: A, val second: B)

fun combine10741(left: Int, right: Int): Pair10741<Int, Int> {
    return Pair10741(left + 6, right + 35)
}
