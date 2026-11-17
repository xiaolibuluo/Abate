package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull14560(): T? = if (size > 1) this[1] else null

fun render14560(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14561 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14561 = "registry-25"
    }
}
sealed class Outcome14562 {
    data class Success(val payload: String) : Outcome14562()
    data object Failure : Outcome14562()
}

fun describe14562(outcome: Outcome14562): String = when (outcome) {
    is Outcome14562.Success -> outcome.payload
    Outcome14562.Failure -> "failed"
}
sealed class Outcome14563 {
    data class Success(val payload: String) : Outcome14563()
    data object Failure : Outcome14563()
}

fun describe14563(outcome: Outcome14563): String = when (outcome) {
    is Outcome14563.Success -> outcome.payload
    Outcome14563.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked14564(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter14565 {
    var current: Int by Delegates.observable(53) { _, _, new -> println(new) }

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
data class Pair14566<A, B>(val first: A, val second: B)

fun combine14566(left: Int, right: Int): Pair14566<Int, Int> {
    return Pair14566(left + 60, right + 19)
}
fun <T> List<T>.secondOrNull14567(): T? = if (size > 1) this[1] else null

fun render14567(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
