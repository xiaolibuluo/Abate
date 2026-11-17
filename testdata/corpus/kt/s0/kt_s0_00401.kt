package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter3290 {
    var current: Int by Delegates.observable(14) { _, _, new -> println(new) }

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
class Counter3291 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
sealed class Outcome3292 {
    data class Success(val payload: String) : Outcome3292()
    data object Failure : Outcome3292()
}

fun describe3292(outcome: Outcome3292): String = when (outcome) {
    is Outcome3292.Success -> outcome.payload
    Outcome3292.Failure -> "failed"
}
data class Pair3293<A, B>(val first: A, val second: B)

fun combine3293(left: Int, right: Int): Pair3293<Int, Int> {
    return Pair3293(left + 35, right + 27)
}
fun <T> List<T>.secondOrNull3294(): T? = if (size > 1) this[1] else null

fun render3294(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome3295 {
    data class Success(val payload: String) : Outcome3295()
    data object Failure : Outcome3295()
}

fun describe3295(outcome: Outcome3295): String = when (outcome) {
    is Outcome3295.Success -> outcome.payload
    Outcome3295.Failure -> "failed"
}
