package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step2290")
fun step2290(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
data class Pair2291<A, B>(val first: A, val second: B)

fun combine2291(left: Int, right: Int): Pair2291<Int, Int> {
    return Pair2291(left + 79, right + 37)
}
object Registry2292 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2292 = "registry-86"
    }
}
fun <T, R> List<T>.mapChecked2293(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step2294")
fun step2294(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
object Registry2295 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2295 = "registry-10"
    }
}
object Registry2296 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2296 = "registry-17"
    }
}
fun <T> List<T>.secondOrNull2297(): T? = if (size > 1) this[1] else null

fun render2297(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
