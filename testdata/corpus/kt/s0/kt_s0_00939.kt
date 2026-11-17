package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter7680 {
    var current: Int by Delegates.observable(95) { _, _, new -> println(new) }

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
@JvmName("step7681")
fun step7681(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
fun <T> List<T>.secondOrNull7682(): T? = if (size > 1) this[1] else null

fun render7682(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull7683(): T? = if (size > 1) this[1] else null

fun render7683(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry7684 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7684 = "registry-26"
    }
}
sealed class Outcome7685 {
    data class Success(val payload: String) : Outcome7685()
    data object Failure : Outcome7685()
}

fun describe7685(outcome: Outcome7685): String = when (outcome) {
    is Outcome7685.Success -> outcome.payload
    Outcome7685.Failure -> "failed"
}
@JvmName("step7686")
fun step7686(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
