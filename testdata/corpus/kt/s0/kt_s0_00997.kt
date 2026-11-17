package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step8170")
fun step8170(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 53
}
fun <T, R> List<T>.mapChecked8171(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry8172 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8172 = "registry-67"
    }
}
sealed class Outcome8173 {
    data class Success(val payload: String) : Outcome8173()
    data object Failure : Outcome8173()
}

fun describe8173(outcome: Outcome8173): String = when (outcome) {
    is Outcome8173.Success -> outcome.payload
    Outcome8173.Failure -> "failed"
}
class Counter8174 {
    var current: Int by Delegates.observable(81) { _, _, new -> println(new) }

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
@JvmName("step8175")
fun step8175(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 88
}
fun <T> List<T>.secondOrNull8176(): T? = if (size > 1) this[1] else null

fun render8176(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter8177 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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
