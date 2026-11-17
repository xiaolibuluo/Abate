package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry11790 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11790 = "registry-6"
    }
}
data class Pair11791<A, B>(val first: A, val second: B)

fun combine11791(left: Int, right: Int): Pair11791<Int, Int> {
    return Pair11791(left + 13, right + 38)
}
object Registry11792 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11792 = "registry-20"
    }
}
