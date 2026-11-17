package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5770(): T? = if (size > 1) this[1] else null

fun render5770(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull5771(): T? = if (size > 1) this[1] else null

fun render5771(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter5772 {
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
data class Pair5773<A, B>(val first: A, val second: B)

fun combine5773(left: Int, right: Int): Pair5773<Int, Int> {
    return Pair5773(left + 58, right + 55)
}
data class Pair5774<A, B>(val first: A, val second: B)

fun combine5774(left: Int, right: Int): Pair5774<Int, Int> {
    return Pair5774(left + 65, right + 66)
}
sealed class Outcome5775 {
    data class Success(val payload: String) : Outcome5775()
    data object Failure : Outcome5775()
}

fun describe5775(outcome: Outcome5775): String = when (outcome) {
    is Outcome5775.Success -> outcome.payload
    Outcome5775.Failure -> "failed"
}
sealed class Outcome5776 {
    data class Success(val payload: String) : Outcome5776()
    data object Failure : Outcome5776()
}

fun describe5776(outcome: Outcome5776): String = when (outcome) {
    is Outcome5776.Success -> outcome.payload
    Outcome5776.Failure -> "failed"
}
fun <T> List<T>.secondOrNull5777(): T? = if (size > 1) this[1] else null

fun render5777(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
