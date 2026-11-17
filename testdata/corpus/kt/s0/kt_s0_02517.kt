package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20690 {
    data class Success(val payload: String) : Outcome20690()
    data object Failure : Outcome20690()
}

fun describe20690(outcome: Outcome20690): String = when (outcome) {
    is Outcome20690.Success -> outcome.payload
    Outcome20690.Failure -> "failed"
}
fun <T> List<T>.secondOrNull20691(): T? = if (size > 1) this[1] else null

fun render20691(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step20692")
fun step20692(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
fun <T, R> List<T>.mapChecked20693(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter20694 {
    var current: Int by Delegates.observable(61) { _, _, new -> println(new) }

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
