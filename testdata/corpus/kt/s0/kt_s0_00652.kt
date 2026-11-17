package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter5340 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
class Counter5341 {
    var current: Int by Delegates.observable(67) { _, _, new -> println(new) }

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
sealed class Outcome5342 {
    data class Success(val payload: String) : Outcome5342()
    data object Failure : Outcome5342()
}

fun describe5342(outcome: Outcome5342): String = when (outcome) {
    is Outcome5342.Success -> outcome.payload
    Outcome5342.Failure -> "failed"
}
sealed class Outcome5343 {
    data class Success(val payload: String) : Outcome5343()
    data object Failure : Outcome5343()
}

fun describe5343(outcome: Outcome5343): String = when (outcome) {
    is Outcome5343.Success -> outcome.payload
    Outcome5343.Failure -> "failed"
}
sealed class Outcome5344 {
    data class Success(val payload: String) : Outcome5344()
    data object Failure : Outcome5344()
}

fun describe5344(outcome: Outcome5344): String = when (outcome) {
    is Outcome5344.Success -> outcome.payload
    Outcome5344.Failure -> "failed"
}
data class Pair5345<A, B>(val first: A, val second: B)

fun combine5345(left: Int, right: Int): Pair5345<Int, Int> {
    return Pair5345(left + 95, right + 20)
}
fun <T> List<T>.secondOrNull5346(): T? = if (size > 1) this[1] else null

fun render5346(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
