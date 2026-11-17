package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6750(): T? = if (size > 1) this[1] else null

fun render6750(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull6751(): T? = if (size > 1) this[1] else null

fun render6751(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair6752<A, B>(val first: A, val second: B)

fun combine6752(left: Int, right: Int): Pair6752<Int, Int> {
    return Pair6752(left + 64, right + 59)
}
fun <T> List<T>.secondOrNull6753(): T? = if (size > 1) this[1] else null

fun render6753(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6754 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6754 = "registry-78"
    }
}
object Registry6755 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6755 = "registry-85"
    }
}
object Registry6756 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6756 = "registry-92"
    }
}
fun <T, R> List<T>.mapChecked6757(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry6758 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6758 = "registry-9"
    }
}
