package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked12420(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step12421")
fun step12421(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
data class Pair12422<A, B>(val first: A, val second: B)

fun combine12422(left: Int, right: Int): Pair12422<Int, Int> {
    return Pair12422(left + 63, right + 63)
}
object Registry12423 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12423 = "registry-70"
    }
}
fun <T> List<T>.secondOrNull12424(): T? = if (size > 1) this[1] else null

fun render12424(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry12425 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12425 = "registry-84"
    }
}
