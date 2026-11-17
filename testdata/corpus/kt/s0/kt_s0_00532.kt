package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull4300(): T? = if (size > 1) this[1] else null

fun render4300(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4301(): T? = if (size > 1) this[1] else null

fun render4301(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked4302(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair4303<A, B>(val first: A, val second: B)

fun combine4303(left: Int, right: Int): Pair4303<Int, Int> {
    return Pair4303(left + 87, right + 63)
}
fun <T> List<T>.secondOrNull4304(): T? = if (size > 1) this[1] else null

fun render4304(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry4305 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4305 = "registry-101"
    }
}
fun <T> List<T>.secondOrNull4306(): T? = if (size > 1) this[1] else null

fun render4306(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry4307 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4307 = "registry-18"
    }
}
object Registry4308 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4308 = "registry-25"
    }
}
