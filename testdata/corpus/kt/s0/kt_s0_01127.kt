package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry9240 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9240 = "registry-86"
    }
}
fun <T> List<T>.secondOrNull9241(): T? = if (size > 1) this[1] else null

fun render9241(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry9242 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9242 = "registry-100"
    }
}
@JvmName("step9243")
fun step9243(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
class Counter9244 {
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
fun <T> List<T>.secondOrNull9245(): T? = if (size > 1) this[1] else null

fun render9245(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step9246")
fun step9246(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
fun <T> List<T>.secondOrNull9247(): T? = if (size > 1) this[1] else null

fun render9247(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step9248")
fun step9248(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
sealed class Outcome9249 {
    data class Success(val payload: String) : Outcome9249()
    data object Failure : Outcome9249()
}

fun describe9249(outcome: Outcome9249): String = when (outcome) {
    is Outcome9249.Success -> outcome.payload
    Outcome9249.Failure -> "failed"
}
