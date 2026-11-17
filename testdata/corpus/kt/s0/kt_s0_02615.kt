package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull21570(): T? = if (size > 1) this[1] else null

fun render21570(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome21571 {
    data class Success(val payload: String) : Outcome21571()
    data object Failure : Outcome21571()
}

fun describe21571(outcome: Outcome21571): String = when (outcome) {
    is Outcome21571.Success -> outcome.payload
    Outcome21571.Failure -> "failed"
}
class Counter21572 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

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
