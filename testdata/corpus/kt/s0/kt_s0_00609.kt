package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4970<A, B>(val first: A, val second: B)

fun combine4970(left: Int, right: Int): Pair4970<Int, Int> {
    return Pair4970(left + 64, right + 11)
}
object Registry4971 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4971 = "registry-71"
    }
}
data class Pair4972<A, B>(val first: A, val second: B)

fun combine4972(left: Int, right: Int): Pair4972<Int, Int> {
    return Pair4972(left + 78, right + 33)
}
