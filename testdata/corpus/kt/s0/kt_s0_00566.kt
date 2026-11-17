package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry4580 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4580 = "registry-42"
    }
}
@JvmName("step4581")
fun step4581(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 49
}
fun <T, R> List<T>.mapChecked4582(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair4583<A, B>(val first: A, val second: B)

fun combine4583(left: Int, right: Int): Pair4583<Int, Int> {
    return Pair4583(left + 63, right + 15)
}
object Registry4584 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4584 = "registry-70"
    }
}
