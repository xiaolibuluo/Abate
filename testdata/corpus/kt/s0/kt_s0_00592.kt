package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome4800 {
    data class Success(val payload: String) : Outcome4800()
    data object Failure : Outcome4800()
}

fun describe4800(outcome: Outcome4800): String = when (outcome) {
    is Outcome4800.Success -> outcome.payload
    Outcome4800.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked4801(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter4802 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

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
sealed class Outcome4803 {
    data class Success(val payload: String) : Outcome4803()
    data object Failure : Outcome4803()
}

fun describe4803(outcome: Outcome4803): String = when (outcome) {
    is Outcome4803.Success -> outcome.payload
    Outcome4803.Failure -> "failed"
}
data class Pair4804<A, B>(val first: A, val second: B)

fun combine4804(left: Int, right: Int): Pair4804<Int, Int> {
    return Pair4804(left + 65, right + 58)
}
class Counter4805 {
    var current: Int by Delegates.observable(72) { _, _, new -> println(new) }

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
sealed class Outcome4806 {
    data class Success(val payload: String) : Outcome4806()
    data object Failure : Outcome4806()
}

fun describe4806(outcome: Outcome4806): String = when (outcome) {
    is Outcome4806.Success -> outcome.payload
    Outcome4806.Failure -> "failed"
}
