package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull19350(): T? = if (size > 1) this[1] else null

fun render19350(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull19351(): T? = if (size > 1) this[1] else null

fun render19351(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked19352(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome19353 {
    data class Success(val payload: String) : Outcome19353()
    data object Failure : Outcome19353()
}

fun describe19353(outcome: Outcome19353): String = when (outcome) {
    is Outcome19353.Success -> outcome.payload
    Outcome19353.Failure -> "failed"
}
class Counter19354 {
    var current: Int by Delegates.observable(65) { _, _, new -> println(new) }

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
data class Pair19355<A, B>(val first: A, val second: B)

fun combine19355(left: Int, right: Int): Pair19355<Int, Int> {
    return Pair19355(left + 72, right + 67)
}
object Registry19356 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19356 = "registry-79"
    }
}
