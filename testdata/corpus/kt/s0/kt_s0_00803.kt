package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6610(): T? = if (size > 1) this[1] else null

fun render6610(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6611 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6611 = "registry-69"
    }
}
data class Pair6612<A, B>(val first: A, val second: B)

fun combine6612(left: Int, right: Int): Pair6612<Int, Int> {
    return Pair6612(left + 76, right + 22)
}
@JvmName("step6613")
fun step6613(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
class Counter6614 {
    var current: Int by Delegates.observable(90) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked6615(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
