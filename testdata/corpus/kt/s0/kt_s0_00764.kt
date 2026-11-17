package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6310")
fun step6310(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
sealed class Outcome6311 {
    data class Success(val payload: String) : Outcome6311()
    data object Failure : Outcome6311()
}

fun describe6311(outcome: Outcome6311): String = when (outcome) {
    is Outcome6311.Success -> outcome.payload
    Outcome6311.Failure -> "failed"
}
fun <T> List<T>.secondOrNull6312(): T? = if (size > 1) this[1] else null

fun render6312(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked6313(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked6314(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome6315 {
    data class Success(val payload: String) : Outcome6315()
    data object Failure : Outcome6315()
}

fun describe6315(outcome: Outcome6315): String = when (outcome) {
    is Outcome6315.Success -> outcome.payload
    Outcome6315.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked6316(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step6317")
fun step6317(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
@JvmName("step6318")
fun step6318(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
class Counter6319 {
    var current: Int by Delegates.observable(26) { _, _, new -> println(new) }

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
