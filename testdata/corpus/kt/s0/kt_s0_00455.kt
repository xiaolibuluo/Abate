package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3720<A, B>(val first: A, val second: B)

fun combine3720(left: Int, right: Int): Pair3720<Int, Int> {
    return Pair3720(left + 88, right + 51)
}
sealed class Outcome3721 {
    data class Success(val payload: String) : Outcome3721()
    data object Failure : Outcome3721()
}

fun describe3721(outcome: Outcome3721): String = when (outcome) {
    is Outcome3721.Success -> outcome.payload
    Outcome3721.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked3722(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter3723 {
    var current: Int by Delegates.observable(12) { _, _, new -> println(new) }

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
sealed class Outcome3724 {
    data class Success(val payload: String) : Outcome3724()
    data object Failure : Outcome3724()
}

fun describe3724(outcome: Outcome3724): String = when (outcome) {
    is Outcome3724.Success -> outcome.payload
    Outcome3724.Failure -> "failed"
}
fun <T> List<T>.secondOrNull3725(): T? = if (size > 1) this[1] else null

fun render3725(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
