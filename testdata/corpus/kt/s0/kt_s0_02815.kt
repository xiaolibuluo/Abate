package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry23220 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23220 = "registry-24"
    }
}
@JvmName("step23221")
fun step23221(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
object Registry23222 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23222 = "registry-38"
    }
}
fun <T, R> List<T>.mapChecked23223(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked23224(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step23225")
fun step23225(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
object Registry23226 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23226 = "registry-66"
    }
}
data class Pair23227<A, B>(val first: A, val second: B)

fun combine23227(left: Int, right: Int): Pair23227<Int, Int> {
    return Pair23227(left + 73, right + 53)
}
