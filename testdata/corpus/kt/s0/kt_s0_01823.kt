package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14990")
fun step14990(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
@JvmName("step14991")
fun step14991(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
@JvmName("step14992")
fun step14992(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
class Counter14993 {
    var current: Int by Delegates.observable(16) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull14994(): T? = if (size > 1) this[1] else null

fun render14994(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step14995")
fun step14995(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
fun <T> List<T>.secondOrNull14996(): T? = if (size > 1) this[1] else null

fun render14996(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome14997 {
    data class Success(val payload: String) : Outcome14997()
    data object Failure : Outcome14997()
}

fun describe14997(outcome: Outcome14997): String = when (outcome) {
    is Outcome14997.Success -> outcome.payload
    Outcome14997.Failure -> "failed"
}
