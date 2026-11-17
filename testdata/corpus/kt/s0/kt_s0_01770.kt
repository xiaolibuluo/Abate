package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14520 {
    data class Success(val payload: String) : Outcome14520()
    data object Failure : Outcome14520()
}

fun describe14520(outcome: Outcome14520): String = when (outcome) {
    is Outcome14520.Success -> outcome.payload
    Outcome14520.Failure -> "failed"
}
sealed class Outcome14521 {
    data class Success(val payload: String) : Outcome14521()
    data object Failure : Outcome14521()
}

fun describe14521(outcome: Outcome14521): String = when (outcome) {
    is Outcome14521.Success -> outcome.payload
    Outcome14521.Failure -> "failed"
}
object Registry14522 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14522 = "registry-77"
    }
}
class Counter14523 {
    var current: Int by Delegates.observable(84) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked14524(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull14525(): T? = if (size > 1) this[1] else null

fun render14525(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14526 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14526 = "registry-8"
    }
}
