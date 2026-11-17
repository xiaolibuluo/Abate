package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull10180(): T? = if (size > 1) this[1] else null

fun render10180(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome10181 {
    data class Success(val payload: String) : Outcome10181()
    data object Failure : Outcome10181()
}

fun describe10181(outcome: Outcome10181): String = when (outcome) {
    is Outcome10181.Success -> outcome.payload
    Outcome10181.Failure -> "failed"
}
class Counter10182 {
    var current: Int by Delegates.observable(61) { _, _, new -> println(new) }

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
object Registry10183 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10183 = "registry-68"
    }
}
fun <T, R> List<T>.mapChecked10184(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair10185<A, B>(val first: A, val second: B)

fun combine10185(left: Int, right: Int): Pair10185<Int, Int> {
    return Pair10185(left + 82, right + 53)
}
sealed class Outcome10186 {
    data class Success(val payload: String) : Outcome10186()
    data object Failure : Outcome10186()
}

fun describe10186(outcome: Outcome10186): String = when (outcome) {
    is Outcome10186.Success -> outcome.payload
    Outcome10186.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked10187(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
