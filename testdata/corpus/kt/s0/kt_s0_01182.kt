package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull9720(): T? = if (size > 1) this[1] else null

fun render9720(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull9721(): T? = if (size > 1) this[1] else null

fun render9721(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter9722 {
    var current: Int by Delegates.observable(45) { _, _, new -> println(new) }

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
class Counter9723 {
    var current: Int by Delegates.observable(52) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked9724(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair9725<A, B>(val first: A, val second: B)

fun combine9725(left: Int, right: Int): Pair9725<Int, Int> {
    return Pair9725(left + 66, right + 36)
}
sealed class Outcome9726 {
    data class Success(val payload: String) : Outcome9726()
    data object Failure : Outcome9726()
}

fun describe9726(outcome: Outcome9726): String = when (outcome) {
    is Outcome9726.Success -> outcome.payload
    Outcome9726.Failure -> "failed"
}
sealed class Outcome9727 {
    data class Success(val payload: String) : Outcome9727()
    data object Failure : Outcome9727()
}

fun describe9727(outcome: Outcome9727): String = when (outcome) {
    is Outcome9727.Success -> outcome.payload
    Outcome9727.Failure -> "failed"
}
