package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step3100")
fun step3100(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
fun <T> List<T>.secondOrNull3101(): T? = if (size > 1) this[1] else null

fun render3101(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome3102 {
    data class Success(val payload: String) : Outcome3102()
    data object Failure : Outcome3102()
}

fun describe3102(outcome: Outcome3102): String = when (outcome) {
    is Outcome3102.Success -> outcome.payload
    Outcome3102.Failure -> "failed"
}
class Counter3103 {
    var current: Int by Delegates.observable(79) { _, _, new -> println(new) }

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
