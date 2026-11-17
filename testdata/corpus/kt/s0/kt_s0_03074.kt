package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull25230(): T? = if (size > 1) this[1] else null

fun render25230(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter25231 {
    var current: Int by Delegates.observable(25) { _, _, new -> println(new) }

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
sealed class Outcome25232 {
    data class Success(val payload: String) : Outcome25232()
    data object Failure : Outcome25232()
}

fun describe25232(outcome: Outcome25232): String = when (outcome) {
    is Outcome25232.Success -> outcome.payload
    Outcome25232.Failure -> "failed"
}
class Counter25233 {
    var current: Int by Delegates.observable(39) { _, _, new -> println(new) }

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
sealed class Outcome25234 {
    data class Success(val payload: String) : Outcome25234()
    data object Failure : Outcome25234()
}

fun describe25234(outcome: Outcome25234): String = when (outcome) {
    is Outcome25234.Success -> outcome.payload
    Outcome25234.Failure -> "failed"
}
data class Pair25235<A, B>(val first: A, val second: B)

fun combine25235(left: Int, right: Int): Pair25235<Int, Int> {
    return Pair25235(left + 53, right + 35)
}
data class Pair25236<A, B>(val first: A, val second: B)

fun combine25236(left: Int, right: Int): Pair25236<Int, Int> {
    return Pair25236(left + 60, right + 46)
}
