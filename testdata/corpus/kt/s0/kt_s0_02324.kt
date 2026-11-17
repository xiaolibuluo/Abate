package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry19090 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19090 = "registry-87"
    }
}
object Registry19091 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19091 = "registry-94"
    }
}
data class Pair19092<A, B>(val first: A, val second: B)

fun combine19092(left: Int, right: Int): Pair19092<Int, Int> {
    return Pair19092(left + 101, right + 35)
}
