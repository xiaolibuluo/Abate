package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23750 {
    data class Success(val payload: String) : Outcome23750()
    data object Failure : Outcome23750()
}

fun describe23750(outcome: Outcome23750): String = when (outcome) {
    is Outcome23750.Success -> outcome.payload
    Outcome23750.Failure -> "failed"
}
sealed class Outcome23751 {
    data class Success(val payload: String) : Outcome23751()
    data object Failure : Outcome23751()
}

fun describe23751(outcome: Outcome23751): String = when (outcome) {
    is Outcome23751.Success -> outcome.payload
    Outcome23751.Failure -> "failed"
}
object Registry23752 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23752 = "registry-48"
    }
}
object Registry23753 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23753 = "registry-55"
    }
}
object Registry23754 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23754 = "registry-62"
    }
}
sealed class Outcome23755 {
    data class Success(val payload: String) : Outcome23755()
    data object Failure : Outcome23755()
}

fun describe23755(outcome: Outcome23755): String = when (outcome) {
    is Outcome23755.Success -> outcome.payload
    Outcome23755.Failure -> "failed"
}
object Registry23756 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23756 = "registry-76"
    }
}
