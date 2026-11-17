package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry11020 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11020 = "registry-72"
    }
}
sealed class Outcome11021 {
    data class Success(val payload: String) : Outcome11021()
    data object Failure : Outcome11021()
}

fun describe11021(outcome: Outcome11021): String = when (outcome) {
    is Outcome11021.Success -> outcome.payload
    Outcome11021.Failure -> "failed"
}
sealed class Outcome11022 {
    data class Success(val payload: String) : Outcome11022()
    data object Failure : Outcome11022()
}

fun describe11022(outcome: Outcome11022): String = when (outcome) {
    is Outcome11022.Success -> outcome.payload
    Outcome11022.Failure -> "failed"
}
object Registry11023 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11023 = "registry-93"
    }
}
class Counter11024 {
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
