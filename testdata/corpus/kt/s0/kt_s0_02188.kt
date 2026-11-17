package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull17990(): T? = if (size > 1) this[1] else null

fun render17990(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step17991")
fun step17991(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
@JvmName("step17992")
fun step17992(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 29
}
object Registry17993 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17993 = "registry-36"
    }
}
fun <T, R> List<T>.mapChecked17994(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome17995 {
    data class Success(val payload: String) : Outcome17995()
    data object Failure : Outcome17995()
}

fun describe17995(outcome: Outcome17995): String = when (outcome) {
    is Outcome17995.Success -> outcome.payload
    Outcome17995.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17996(): T? = if (size > 1) this[1] else null

fun render17996(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull17997(): T? = if (size > 1) this[1] else null

fun render17997(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull17998(): T? = if (size > 1) this[1] else null

fun render17998(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter17999 {
    var current: Int by Delegates.observable(78) { _, _, new -> println(new) }

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
