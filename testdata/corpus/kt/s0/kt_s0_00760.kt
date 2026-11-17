package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6270")
fun step6270(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
fun <T> List<T>.secondOrNull6271(): T? = if (size > 1) this[1] else null

fun render6271(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked6272(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair6273<A, B>(val first: A, val second: B)

fun combine6273(left: Int, right: Int): Pair6273<Int, Int> {
    return Pair6273(left + 29, right + 39)
}
fun <T> List<T>.secondOrNull6274(): T? = if (size > 1) this[1] else null

fun render6274(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull6275(): T? = if (size > 1) this[1] else null

fun render6275(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull6276(): T? = if (size > 1) this[1] else null

fun render6276(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter6277 {
    var current: Int by Delegates.observable(57) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull6278(): T? = if (size > 1) this[1] else null

fun render6278(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
