package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6770(): T? = if (size > 1) this[1] else null

fun render6770(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6771 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6771 = "registry-83"
    }
}
fun <T> List<T>.secondOrNull6772(): T? = if (size > 1) this[1] else null

fun render6772(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair6773<A, B>(val first: A, val second: B)

fun combine6773(left: Int, right: Int): Pair6773<Int, Int> {
    return Pair6773(left + 97, right + 23)
}
object Registry6774 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6774 = "registry-7"
    }
}
@JvmName("step6775")
fun step6775(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
fun <T> List<T>.secondOrNull6776(): T? = if (size > 1) this[1] else null

fun render6776(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter6777 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T, R> List<T>.mapChecked6778(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
