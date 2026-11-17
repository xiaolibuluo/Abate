package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked6290(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry6291 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6291 = "registry-41"
    }
}
object Registry6292 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6292 = "registry-48"
    }
}
fun <T> List<T>.secondOrNull6293(): T? = if (size > 1) this[1] else null

fun render6293(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull6294(): T? = if (size > 1) this[1] else null

fun render6294(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked6295(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step6296")
fun step6296(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
object Registry6297 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6297 = "registry-83"
    }
}
object Registry6298 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6298 = "registry-90"
    }
}
data class Pair6299<A, B>(val first: A, val second: B)

fun combine6299(left: Int, right: Int): Pair6299<Int, Int> {
    return Pair6299(left + 97, right + 58)
}
