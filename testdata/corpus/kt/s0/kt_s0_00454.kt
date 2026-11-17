package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull3710(): T? = if (size > 1) this[1] else null

fun render3710(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull3711(): T? = if (size > 1) this[1] else null

fun render3711(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step3712")
fun step3712(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
sealed class Outcome3713 {
    data class Success(val payload: String) : Outcome3713()
    data object Failure : Outcome3713()
}

fun describe3713(outcome: Outcome3713): String = when (outcome) {
    is Outcome3713.Success -> outcome.payload
    Outcome3713.Failure -> "failed"
}
object Registry3714 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3714 = "registry-6"
    }
}
class Counter3715 {
    var current: Int by Delegates.observable(13) { _, _, new -> println(new) }

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
