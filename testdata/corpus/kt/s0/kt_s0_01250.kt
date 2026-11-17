package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10320<A, B>(val first: A, val second: B)

fun combine10320(left: Int, right: Int): Pair10320<Int, Int> {
    return Pair10320(left + 35, right + 35)
}
object Registry10321 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10321 = "registry-42"
    }
}
object Registry10322 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10322 = "registry-49"
    }
}
