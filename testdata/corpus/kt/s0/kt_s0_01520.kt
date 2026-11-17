package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step12590")
fun step12590(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
object Registry12591 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12591 = "registry-83"
    }
}
@JvmName("step12592")
fun step12592(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
@JvmName("step12593")
fun step12593(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
@JvmName("step12594")
fun step12594(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 7
}
sealed class Outcome12595 {
    data class Success(val payload: String) : Outcome12595()
    data object Failure : Outcome12595()
}

fun describe12595(outcome: Outcome12595): String = when (outcome) {
    is Outcome12595.Success -> outcome.payload
    Outcome12595.Failure -> "failed"
}
fun <T> List<T>.secondOrNull12596(): T? = if (size > 1) this[1] else null

fun render12596(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter12597 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked12598(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
