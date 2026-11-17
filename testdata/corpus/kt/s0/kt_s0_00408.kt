package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry3340 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3340 = "registry-79"
    }
}
fun <T> List<T>.secondOrNull3341(): T? = if (size > 1) this[1] else null

fun render3341(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step3342")
fun step3342(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
class Counter3343 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull3344(): T? = if (size > 1) this[1] else null

fun render3344(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter3345 {
    var current: Int by Delegates.observable(17) { _, _, new -> println(new) }

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
sealed class Outcome3346 {
    data class Success(val payload: String) : Outcome3346()
    data object Failure : Outcome3346()
}

fun describe3346(outcome: Outcome3346): String = when (outcome) {
    is Outcome3346.Success -> outcome.payload
    Outcome3346.Failure -> "failed"
}
