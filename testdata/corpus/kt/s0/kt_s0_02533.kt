package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20850 {
    data class Success(val payload: String) : Outcome20850()
    data object Failure : Outcome20850()
}

fun describe20850(outcome: Outcome20850): String = when (outcome) {
    is Outcome20850.Success -> outcome.payload
    Outcome20850.Failure -> "failed"
}
fun <T> List<T>.secondOrNull20851(): T? = if (size > 1) this[1] else null

fun render20851(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked20852(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter20853 {
    var current: Int by Delegates.observable(68) { _, _, new -> println(new) }

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
@JvmName("step20854")
fun step20854(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
sealed class Outcome20855 {
    data class Success(val payload: String) : Outcome20855()
    data object Failure : Outcome20855()
}

fun describe20855(outcome: Outcome20855): String = when (outcome) {
    is Outcome20855.Success -> outcome.payload
    Outcome20855.Failure -> "failed"
}
sealed class Outcome20856 {
    data class Success(val payload: String) : Outcome20856()
    data object Failure : Outcome20856()
}

fun describe20856(outcome: Outcome20856): String = when (outcome) {
    is Outcome20856.Success -> outcome.payload
    Outcome20856.Failure -> "failed"
}
