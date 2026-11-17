package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull2620(): T? = if (size > 1) this[1] else null

fun render2620(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry2621 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2621 = "registry-23"
    }
}
sealed class Outcome2622 {
    data class Success(val payload: String) : Outcome2622()
    data object Failure : Outcome2622()
}

fun describe2622(outcome: Outcome2622): String = when (outcome) {
    is Outcome2622.Success -> outcome.payload
    Outcome2622.Failure -> "failed"
}
class Counter2623 {
    var current: Int by Delegates.observable(37) { _, _, new -> println(new) }

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
