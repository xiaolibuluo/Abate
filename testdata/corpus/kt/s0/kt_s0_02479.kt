package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull20400(): T? = if (size > 1) this[1] else null

fun render20400(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter20401 {
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
fun <T> List<T>.secondOrNull20402(): T? = if (size > 1) this[1] else null

fun render20402(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step20403")
fun step20403(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
sealed class Outcome20404 {
    data class Success(val payload: String) : Outcome20404()
    data object Failure : Outcome20404()
}

fun describe20404(outcome: Outcome20404): String = when (outcome) {
    is Outcome20404.Success -> outcome.payload
    Outcome20404.Failure -> "failed"
}
class Counter20405 {
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
@JvmName("step20406")
fun step20406(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 86
}
sealed class Outcome20407 {
    data class Success(val payload: String) : Outcome20407()
    data object Failure : Outcome20407()
}

fun describe20407(outcome: Outcome20407): String = when (outcome) {
    is Outcome20407.Success -> outcome.payload
    Outcome20407.Failure -> "failed"
}
