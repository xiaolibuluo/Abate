package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair6830<A, B>(val first: A, val second: B)

fun combine6830(left: Int, right: Int): Pair6830<Int, Int> {
    return Pair6830(left + 57, right + 32)
}
object Registry6831 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6831 = "registry-64"
    }
}
