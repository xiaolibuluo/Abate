package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair5710<A, B>(val first: A, val second: B)

fun combine5710(left: Int, right: Int): Pair5710<Int, Int> {
    return Pair5710(left + 56, right + 41)
}
class Counter5711 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
data class Pair5712<A, B>(val first: A, val second: B)

fun combine5712(left: Int, right: Int): Pair5712<Int, Int> {
    return Pair5712(left + 70, right + 63)
}
sealed class Outcome5713 {
    data class Success(val payload: String) : Outcome5713()
    data object Failure : Outcome5713()
}

fun describe5713(outcome: Outcome5713): String = when (outcome) {
    is Outcome5713.Success -> outcome.payload
    Outcome5713.Failure -> "failed"
}
fun <T> List<T>.secondOrNull5714(): T? = if (size > 1) this[1] else null

fun render5714(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair5715<A, B>(val first: A, val second: B)

fun combine5715(left: Int, right: Int): Pair5715<Int, Int> {
    return Pair5715(left + 91, right + 35)
}
sealed class Outcome5716 {
    data class Success(val payload: String) : Outcome5716()
    data object Failure : Outcome5716()
}

fun describe5716(outcome: Outcome5716): String = when (outcome) {
    is Outcome5716.Success -> outcome.payload
    Outcome5716.Failure -> "failed"
}
