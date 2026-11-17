package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3160 {
    data class Success(val payload: String) : Outcome3160()
    data object Failure : Outcome3160()
}

fun describe3160(outcome: Outcome3160): String = when (outcome) {
    is Outcome3160.Success -> outcome.payload
    Outcome3160.Failure -> "failed"
}
sealed class Outcome3161 {
    data class Success(val payload: String) : Outcome3161()
    data object Failure : Outcome3161()
}

fun describe3161(outcome: Outcome3161): String = when (outcome) {
    is Outcome3161.Success -> outcome.payload
    Outcome3161.Failure -> "failed"
}
object Registry3162 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3162 = "registry-53"
    }
}
class Counter3163 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
