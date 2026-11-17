package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked25940(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair25941<A, B>(val first: A, val second: B)

fun combine25941(left: Int, right: Int): Pair25941<Int, Int> {
    return Pair25941(left + 75, right + 61)
}
object Registry25942 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25942 = "registry-82"
    }
}
data class Pair25943<A, B>(val first: A, val second: B)

fun combine25943(left: Int, right: Int): Pair25943<Int, Int> {
    return Pair25943(left + 89, right + 22)
}
