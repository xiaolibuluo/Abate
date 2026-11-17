package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6250(): T? = if (size > 1) this[1] else null

fun render6250(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6251 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6251 = "registry-86"
    }
}
fun <T> List<T>.secondOrNull6252(): T? = if (size > 1) this[1] else null

fun render6252(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter6253 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull6254(): T? = if (size > 1) this[1] else null

fun render6254(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome6255 {
    data class Success(val payload: String) : Outcome6255()
    data object Failure : Outcome6255()
}

fun describe6255(outcome: Outcome6255): String = when (outcome) {
    is Outcome6255.Success -> outcome.payload
    Outcome6255.Failure -> "failed"
}
@JvmName("step6256")
fun step6256(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
fun <T, R> List<T>.mapChecked6257(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
